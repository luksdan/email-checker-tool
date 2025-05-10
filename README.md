# Domain Scanner (MX, SPF, DMARC)

Este é um utilitário em Go que verifica automaticamente se um ou mais domínios possuem registros de **MX**, **SPF** e **DMARC** e exporta os resultados em um arquivo CSV.

## Funcionalidades

- Verifica se o domínio possui registro **MX** (Mail Exchange)
- Identifica se o domínio possui registro **SPF** (Sender Policy Framework) e captura o valor
- Identifica se o domínio possui registro **DMARC** e captura o valor
- Exporta os resultados no arquivo `result.csv`

## Requisitos

- Go 1.16 ou superior instalado

## Instalação

Clone o repositório e compile o projeto:

```bash
git clone https://github.com/seu-usuario/domain-scanner.git
cd domain-scanner
go build -o domain-scanner
```

## Como usar
Rode o programa e digite os domínios (um por linha). Finalize com Ctrl+D (Linux/Mac) ou Ctrl+Z + Enter (Windows):
```bash
./domain-scanner
google.com
example.com
seu-dominio.com
```

## Exemplo de saída (result.csv)
domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord
google.com,true,true,v=spf1 include:_spf.google.com ~all,true,v=DMARC1; p=none; rua=mailto:dmarc-reports@google.com
example.com,false,false,,false,