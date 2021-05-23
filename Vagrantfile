# coding: utf-8
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/bionic64"

  # Provisiona a máquina para o servidor da aplicação
  config.vm.define "application" do |application_config|
    application_config.vm.hostname = "application"
    application_config.vm.network :private_network,
                          :ip => "192.168.33.12"
    # Roda o playbook de instalação e configuração do servidor web
    config.vm.provision "ansible" do |ansible|
      ansible.verbose = "v"
      ansible.playbook = "playbooks/application.yml"
    end
  end

  # Provisiona a máquina para o servidor do Jenkins
  config.vm.define "jenkins" do |jenkins_config|
    jenkins_config.vm.box = "debian/buster64"
    jenkins_config.vm.hostname = "jenkins"
    jenkins_config.vm.network :private_network,
                          :ip => "192.168.33.13"
    # Roda o playbook de instalação e configuração do servidor do Jenkins
    config.vm.provision "ansible" do |ansible|
      ansible.verbose = "v"
      ansible.playbook = "playbooks/jenkins.yml"
    end
  end

  # Provisiona a máquina para o servidor de monitoração
  config.vm.define "monitor" do |monitor_config|
    monitor_config.vm.hostname = "monitor"
    monitor_config.vm.network :private_network,
                              :ip => "192.168.33.14"
    # Roda o playbook de instalação e configuração do servidor de monitoração
    # instalando e configurando o nagios3
    config.vm.provision "ansible" do |ansible|
      ansible.verbose = "v"
      ansible.playbook = "playbooks/monitor.yml"
    end
  end
end
