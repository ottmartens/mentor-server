pipeline {
    agent any {
        withEnv(["GOROOT=${goRoot}","GOPATH=${goPath}", "PATH=${goRoot}/bin:${goPath}/bin:${path}"]) {            
            stages {
            
                goRoot='/usr/local/go'
                goPath='/root/go'
                path = env.PATH

                stage('verifying environment') {    
                    sh 'go version'
                }
                
                stage('pull changes') {
                    git 'https://github.com/ottmartens/mentor-server'
                }
                
                stage('tests') {
                    echo 'no tests configured!'
                }
                
                stage('build') {
                    sh 'go build -o mentor-server'
                    sh 'ls'
                }
                
                stage('deploy & reload service') {
                    sh '/bin/systemctl stop mentor-server.service'
                    sh 'cp mentor-server /var/www/mentor-server'
                    sh '/bin/systemctl start mentor-server.service'
                }
            }
        }
    }
}

