# Weather API with CEP

Este é um sistema em Go que recebe um CEP válido, identifica a cidade associada e retorna o clima atual em Celsius, Fahrenheit e Kelvin. O sistema é implantado no Google Cloud Run e pode ser testado localmente usando Docker e Docker Compose.

## Requisitos

- Go
- Docker e Docker Compose
- Conta no Google Cloud (para deploy no Cloud Run)
- Chaves de API para:
  - [WeatherAPI](https://www.weatherapi.com/) (consulta de clima)

## Como Funciona

1. O sistema recebe um CEP válido de 8 dígitos.
2. Consulta a localização (cidade e estado) usando a API ViaCEP.
3. Consulta o clima atual da localização usando a API WeatherAPI.
4. Retorna as temperaturas em Celsius, Fahrenheit e Kelvin.

### Exemplo de Resposta

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

## Configuração

1. Clone o repositório
```
git clone https://github.com/paulagates/cloud-run.git
cd cloud-run

```
2. Configure as variáveis de ambiente
Crie um arquivo .env na raiz do projeto com as seguintes variáveis:

```
WEATHER_API_KEY=sua_chave_da_weatherapi

```

## Acesso ao sistema

Acesso pela URL: https://cloud-run-305796103700.southamerica-east1.run.app

# Rotas

```
GET /weather?cep=XXXXX-XXX
```
Retorna o clima atual para o CEP fornecido.

Exemplo: GET https://cloud-run-305796103700.southamerica-east1.run.app/weather?cep=01001000

## Testes Automatizados

Para rodar os testes automatizados, use o Docker Compose:

```
docker-compose run --rm tester

```

Os testes cobrem os seguintes cenários:

- No handler: Localidade do CEP válido e inválido, buscar o weather e o status esperado.

- No service: Localidade do CEP válido e inválido, buscar o weather.