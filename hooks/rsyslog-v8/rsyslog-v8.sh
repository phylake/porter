# (c) 2017 Adobe. All rights reserved.
# This file is licensed to you under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License. You may obtain a copy
# of the License at http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed under
# the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
# OF ANY KIND, either express or implied. See the License for the specific language
# governing permissions and limitations under the License.

#!/bin/bash -e
cat > /etc/yum.repos.d/rsyslog.repo <<EOF
[rsyslog_v8]
name=Adiscon CentOS-7 - local packages for x86_64
baseurl=http://rpms.adiscon.com/v8-stable/epel-7/x86_64
enabled=1
gpgcheck=0
gpgkey=http://rpms.adiscon.com/RPM-GPG-KEY-Adiscon
protect=1
EOF

yum -y install \
libtool \
libgcrypt-devel \
libuuid-devel \
zlib-devel \
json-c-devel \
libestr-devel \
liblogging-devel \
libfastjson4-devel \
byacc

cd /home/ec2-user

curl --retry 3 -sL http://www.rsyslog.com/files/download/rsyslog/rsyslog-8.25.0.tar.gz | tar -xz

cd /home/ec2-user/rsyslog-8.25.0

./configure --prefix / 1>/dev/null
make -j2 1>/dev/null

# wait for success of the above steps before stopping rsyslog
service rsyslog stop

# this file contains flags that break v8
rm -f /etc/sysconfig/rsyslog

# install
make install 1>/dev/null

# updated config
cat > /etc/rsyslog.conf <<EOF
module(load="imuxsock")
module(load="imtcp" threads="2")
module(load="imudp" threads="2")

main_queue(
  queue.size="1000000"
  queue.debatchsize="1000"
  queue.workerthreads="2"
)

input(type="imtcp" port="514" ruleset="logFile")
input(type="imudp" port="514" ruleset="logFile")

ruleset(name="logFile"
        queue.type="fixedArray"
        queue.size="250000"
        queue.dequeueBatchSize="4096"
        queue.workerThreads="4"
        queue.workerThreadMinimumMessages="60000"
       ) {
    action(type="omfile" file="/var/log/porter.log"
           ioBufferSize="64k" flushOnTXEnd="off"
           asyncWriting="on")
}
EOF

# start
service rsyslog start
