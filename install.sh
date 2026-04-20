#!/bin/sh
set -e

REPO="matheusbrantdev/ghswitch"
BIN="/usr/local/bin/ghswitch"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
  x86_64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

LATEST=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | cut -d'"' -f4)

if [ -z "$LATEST" ]; then
  echo "Could not fetch latest release."
  exit 1
fi

URL="https://github.com/$REPO/releases/download/$LATEST/ghswitch_${OS}_${ARCH}.tar.gz"

echo "Installing ghswitch $LATEST..."

curl -sL "$URL" | tar -xz -C /tmp ghswitch
chmod +x /tmp/ghswitch
mv /tmp/ghswitch "$BIN"

echo "ghswitch installed at $BIN"

# configure bash prompt
BASHRC="$HOME/.bashrc"

if ! grep -q "ghswitch_prompt" "$BASHRC" 2>/dev/null; then
  cat >> "$BASHRC" << 'EOF'

ghswitch_prompt() {
  local active=$(cat ~/.ghswitch/active 2>/dev/null)
  [ -n "$active" ] && echo -e "\033[2m(\033[0m\033[1;36mgh\033[0m\033[2m·\033[0m\033[1;37m$active\033[0m\033[2m)\033[0m "
}

PS1='$(ghswitch_prompt)'"$PS1"
EOF
  echo "Prompt configured in $BASHRC"
  echo "Run: source ~/.bashrc"
fi

echo "Done. Run 'ghswitch add' to create your first profile."
