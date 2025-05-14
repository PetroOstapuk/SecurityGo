# SecurityGo

## Port scanner
```
 go build -o portscanner port_scanner.go
```
 Start scanning
```
 ./portscanner
```

Result:
```
🖥️ Введіть хост для сканування: scanme.nmap.org
🔢 Введіть кількість портів для сканування (наприклад, 1024): 64000
🚀 Сканування портів...

📋 Відкриті порти:
  ✅  22
  ✅  80
  ✅  9929
  ✅  31337

⏳ Час: 2m2s

✅ Сканування завершено.

```