node {

    stage('pull changes') {
        git 'https://github.com/ottmartens/mentor-server'
    }
        
    stage('tests') {
        echo 'no tests configured!'
    }
    
    stage('build Docker image') {
        sh 'docker build -t mentor-server .'
    }
    
    stage('deploy & reload service') {
        sh 'docker stack deploy -c docker-compose.yml mentor-server'
    }
}