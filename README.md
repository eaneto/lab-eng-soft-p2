# P2

![Visão Geral da Arquitetura](./arquitetura.png)

## Set Up

Se você quiser rodar a aplicação de teste é necessário configurar as variáveis
de ambiente da AWS.

```bash
export AWS_ACCOUNT=sua-conta
export AWS_ACCESS_KEY_ID=secret-key-id
export AWS_SECRET_ACCESS_KEY=secret-access-key
```

## Ansible

Para instalar o jenkins no servidor é preciso da role customizada do jenkins.

```bash
ansible-galaxy install emmetog.jenkins
```

## Terraform

O arquivo do terraform é utilizado para criar a fila e o tópico de
notificação.

```bash
cd terraform/
terraform init
terraform apply -auto-approve
```

## Vagrant

```bash
vagrant box add ubuntu/bionic64
vagrant box add debian/buster64
vagrant up
```
