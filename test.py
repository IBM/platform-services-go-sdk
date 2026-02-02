#!/usr/bin/env python3
"""
bkill-switch-script.py – Scan the npm ecosystem for “bad” modules.

Features
--------
* Downloads a CSV of bad packages from a Google-Sheet export
  (fallback to a local CSV if the download fails).
* Loads the bad-list into a set of keys:
    – “name@version” when a version is specified
    – “name” when the CSV leaves the version empty.
* Inspects the global npm tree (`npm list -g --depth=0 --json`)
  and the local project tree (`npm list --depth=0 --json`).
* If `package-lock.json` is missing, falls back to a recursive
  `npm ls --depth=Infinity --json` to cover transitive deps.
* Exits with status 1 if a bad package is found, otherwise
  prints a success message.

Dependencies
------------
* Python 3.8+ (uses `urllib.request`, `csv`, `json`, `subprocess`).
* `npm` must be available on the PATH.
"""

import json
import os
import subprocess
import sys
import urllib.request
import csv
from pathlib import Path
from typing import Dict, Set, Tuple, List, Optional
import ssl

# Globally disable SSL certificate verification
ssl._create_default_https_context = ssl._create_unverified_context

# --------------------------------------------------------------------
# Configuration
# --------------------------------------------------------------------
BAD_PACKAGES_URL = (
    "https://docs.google.com/spreadsheets/d/1ma78Oc9mCNZlablv5x9xjKTziOmWHUwL6jHIcBBwlyA/export?format=csv&gid=0"
)
LOCAL_BAD_CSV = Path("bad-packages.csv")   # fallback if download fails

# --------------------------------------------------------------------
# Helpers
# --------------------------------------------------------------------
def download_csv(url: str, dest: Path) -> bool:
    """Download a CSV from *url* to *dest*.
    Returns True on success, False otherwise."""
    try:
        with urllib.request.urlopen(url, timeout=10) as r, open(dest, "wb") as w:
            w.write(r.read())
        return True
    except Exception as exc:
        print(f"⚠️  Failed to download bad-list from {url} ({exc})", file=sys.stderr)
        return False


def load_bad_list(csv_path: Path) -> Set[str]:
    """Read the CSV and return a set of bad keys."""
    bad: Set[str] = set()
    with csv_path.open(newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        i = 0
        for row in reader:
            pkg = row.get("package name", "").strip()
            ver = row.get("package version(s)", "").strip()
            if not pkg or pkg.lower().startswith("package name"):
                continue
            key = f"{pkg}@{ver}" if ver else pkg
            i += 1
            print(f"     - Loaded bad package: {i:>4} : {key}")
            bad.add(key)
    return bad


def run_npm_list(scope: str) -> Dict:
    """Run `npm list` for the given *scope* ('global' or 'local') and
    return the parsed JSON. Returns {} if it fails gracefully."""
    if scope == "global":
        cmd = ["npm", "list", "-g", "--depth=0", "--json"]
    else:
        # don't use --package-lock-only, it's too strict
        cmd = ["npm", "list", "--depth=0", "--json"]

    try:
        out = subprocess.check_output(cmd, stderr=subprocess.DEVNULL, text=True)
        return json.loads(out)
    except subprocess.CalledProcessError as exc:
        print(f"⚠️  npm list ({scope}) failed: {exc}", file=sys.stderr)
        return {}
    except json.JSONDecodeError as exc:
        print(f"⚠️  Failed to parse npm list output ({scope}): {exc}", file=sys.stderr)
        return {}


def run_npm_ls_infinite() -> Dict:
    """Run `npm ls --depth=Infinity --json` and return the parsed JSON."""
    cmd = ["npm", "ls", "--depth=Infinity", "--json"]
    try:
        out = subprocess.check_output(cmd, stderr=subprocess.DEVNULL, text=True)
        return json.loads(out)
    except subprocess.CalledProcessError as exc:
        print(f"⚠️  npm ls (infinite) failed: {exc}", file=sys.stderr)
        return {}
    except json.JSONDecodeError as exc:
        print(f"⚠️  Failed to parse npm ls output ({exc})", file=sys.stderr)
        return {}


def top_level_dependencies(data: Dict) -> List[Tuple[str, str]]:
    """Return a list of (name, version) tuples for the top-level
    dependencies recorded by `npm list --depth=0`."""
    deps = data.get("dependencies", {})
    return [(name, info.get("version")) for name, info in deps.items() if info.get("version")]


def walk_lockfile(node: Dict, acc: Optional[List[Tuple[str, str]]] = None) -> List[Tuple[str, str]]:
    """Recursively walk a lockfile node and accumulate (name, version)."""
    if acc is None:
        acc = []
    for name, info in node.get("dependencies", {}).items():
        ver = info.get("version")
        if ver:
            acc.append((name, ver))
        # Recurse into nested deps
        if "dependencies" in info:
            walk_lockfile(info, acc)
    return acc


def check_packages(installed: List[Tuple[str, str]], bad_set: Set[str], scope: str) -> None:
    """Print an error and exit if any installed package matches the bad set."""
    for name, ver in installed:
        key = f"{name}@{ver}"
        print(f"     - Checking {key}")
        if key in bad_set or name in bad_set:
            print(f"❌  Bad package detected ({scope}): {key}", file=sys.stderr)
            sys.exit(1)


# --------------------------------------------------------------------
# Main logic
# --------------------------------------------------------------------
def main() -> None:
    # 1️⃣  Obtain the bad-list (download → fallback to local)
    temp_csv: Optional[Path] = None
    if download_csv(BAD_PACKAGES_URL, temp_csv := Path("/tmp/bad-packages.csv")):
        csv_path = temp_csv
    else:
        if not LOCAL_BAD_CSV.is_file():
            print(f"❌  Local CSV {LOCAL_BAD_CSV} not found", file=sys.stderr)
            sys.exit(1)
        csv_path = LOCAL_BAD_CSV

    bad_set = load_bad_list(csv_path)

    # 2️⃣  Scan global tree
    print("🔍  Scanning globally installed npm modules …")
    global_data = run_npm_list("global")
    if global_data:
        check_packages(top_level_dependencies(global_data), bad_set, "global")
    else:
        print("⚠️   Skipping global npm scan (no valid output).", file=sys.stderr)

    # 3️⃣  Scan local tree (project root)
    if not Path("package.json").is_file():
        print("⚠️   No package.json found – skipping local scan.", file=sys.stderr)
    else:
        print("🔍  Scanning locally installed npm modules …")
        local_data = run_npm_list("local")
        if local_data:
            check_packages(top_level_dependencies(local_data), bad_set, "local")
        else:
            print("⚠️   Skipping local npm scan (no valid output).", file=sys.stderr)

        # 4️⃣  Also check the lockfile if present
        lock_path = Path("package-lock.json")
        if lock_path.is_file():
            print("🔎  Scanning package-lock.json for bad packages …")
            try:
                lock_json = json.loads(lock_path.read_text(encoding="utf-8"))
                lock_deps = walk_lockfile(lock_json)
                check_packages(lock_deps, bad_set, "lockfile")
            except Exception as exc:
                print(f"⚠️  Failed to parse package-lock.json ({exc})", file=sys.stderr)
        else:
            print("⚠️   No package-lock.json – walking the whole npm tree instead.")
            lock_json = run_npm_ls_infinite()
            if lock_json:
                lock_deps = walk_lockfile(lock_json)
                check_packages(lock_deps, bad_set, "npm-ls-infinity")

    # 5️⃣  Success
    print("✅  No bad npm packages detected.")
    sys.exit(0)


if __name__ == "__main__":
    main()
