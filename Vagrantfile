Vagrant.configure('2') do |config|
    config.vm.box = 'bento/ubuntu-16.04'
    config.ssh.insert_key = false
    config.vm.define 'gpl' do |t|
        t.vm.hostname = 'gpl'
        t.vm.network(:private_network, ip: '192.168.142.142')
        t.vm.provider 'virtualbox' do |v|
          v.memory = 2048
        end
    end
end
