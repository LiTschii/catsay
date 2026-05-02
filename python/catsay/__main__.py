"""Entry point: python -m catsay"""

import os
import sys

from catsay._binary import ensure_binary


def main() -> None:
    binary = ensure_binary()
    # Replace the current process with the binary (Unix) or spawn it (Windows)
    if sys.platform == "win32":
        import subprocess
        result = subprocess.run([str(binary)] + sys.argv[1:])
        sys.exit(result.returncode)
    else:
        os.execv(str(binary), [str(binary)] + sys.argv[1:])


if __name__ == "__main__":
    main()
