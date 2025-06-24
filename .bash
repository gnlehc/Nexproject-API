# Install Debugger
go install github.com/go-delve/delve/cmd/dlv@latest

# Add to ENV
echo 'export PATH=$(go env GOPATH)/bin:$PATH' >> ~/.zshrc
source ~/.zshrc

# Run in Debug Mode
dlv debug main.go
