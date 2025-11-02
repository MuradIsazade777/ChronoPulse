# Contributing to ChronoPulse

Thank you for considering contributing to ChronoPulse! We welcome pull requests, issues, and suggestions.

## 🧭 How to Contribute
 
1. Fork the repository 
2. Create a new branch (`git checkout -b feature-name`)
3. Commit your changes (`git commit -m 'Add feature'`) 
4. Push to your fork (`git push origin feature-name`)
5. Open a pull request

## ✅ Code Guidelines 

- Use idiomatic Go
- Keep functions modular and testable
- Maintain consistent formatting (`go fmt`)
- Write meaningful commit messages

## 📄 Documentation

If you add new features, update the README and module comments.

## 🧪 Testing

Ensure your code builds and runs cleanly:

```bash
go vet ./...
go build ./...
go run cmd/chronopulse/main.go
```
