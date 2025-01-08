# Notificador (SMS via Amazon SNS)

Este é um microserviço em Go (Golang) que envia notificações via SMS utilizando o Amazon Simple Notification Service (SNS). Ele foi projetado para ser eficiente, simples e escalável, aproveitando as goroutines para gerenciar a concorrência.

## **Recursos**
- Envio de mensagens SMS diretamente para números de telefone.
- Configuração automática de credenciais da AWS via `~/.aws/credentials`.
- Leve e rápido, ideal para integração com outros sistemas ou como microsserviço independente.

## **Requisitos**
- Go 1.18 ou superior.
- Conta na AWS com o serviço Amazon SNS habilitado.
- Credenciais configuradas no arquivo `~/.aws/credentials`.

## **Instalação**
1. Clone este repositório:
   ```bash
   git clone https://github.com/mPadovani74/NotificadorSMS-GOang
   cd NotificadorSMS-GOang
