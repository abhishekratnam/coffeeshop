#!/bin/bash 
 
echo "upgrading existing apparmor dependencies according to the kubearmor." 
sudo apt-get install build-essential
sudo apt-get install --reinstall make
sudo apt-get install manpages-dev
gcc --version
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
echo "$(<kubectl.sha256) kubectl" | sha256sum --check
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl 
git clone git://git.kernel.org/pub/scm/linux/kernel/git/jj/linux-apparmor.git linux-apparmor
cd linux-apparmor
make -C $KDIR M=$PWD
sudo ./build.sh --kernel-dir /lib/modules/$(uname -r)/build
diff --git a/Makefile b/Makefile
index 5c18baa..e342473 100644
--- a/Makefile
+++ b/Makefile
 @@ -612,6 +612,12 @@ endif # $(dot-config)
 # Defaults to vmlinux, but the arch makefile usually adds further targets
 all: vmlinux
 
+# force no-pie for distro compilers that enable pie by default
+KBUILD_CFLAGS += $(call cc-option, -fno-pie)
+KBUILD_CFLAGS += $(call cc-option, -no-pie)
+KBUILD_AFLAGS += $(call cc-option, -fno-pie)
+KBUILD_CPPFLAGS += $(call cc-option, -fno-pie) 
+
 # The arch Makefile can set ARCH_{CPP,A,C}FLAGS to override the default
 # values of the respective KBUILD_* variables
ARCH_CPPFLAGS :=

make oldconfig
make silentoldconfig

make
make install
cd ..
#Kernel config 
var=$(uname -r | cut -d'.' -f1)
if [ $var -gt 4 ]
then
   echo "Kernel Okay"
else
   echo "Kernel compatibility error"
   
   
fi

# checking linux security module (LSM) is available
var=$(cat /sys/kernel/security/apparmor/|grep CONFIG_SECURITY /boot/config-`uname -r`|grep not) 
if [ $var ]
then
   echo "Please check LSM settings"
else
    echo "LSM Okay"
git clone https://github.com/accuknox/KubeArmor.git
cd KubeArmor
cd contributions/bare-metal
./setup.sh
cd microk8s
./install_microk8s.sh
cd ..
cd ..
cd vagrant
./setup.sh
./create.sh
ssh-keygen -> [Enter] -> [Enter] -> [Enter]
vagrant up




