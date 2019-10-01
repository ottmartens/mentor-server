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

    stage('push image to local registry') {
        sh 'docker tag mentor-server localhost:5000/mentor-server-local'
        sh 'docker push localhost:5000/mentor-server-local'
    }

    stage('remove old local images') {
        sh 'docker image remove mentor-server'
        sh 'docker image remove localhost:5000/mentor-server-local'
    }
    
    stage('deploy & reload service') {
        sh 'docker stack deploy -c docker-compose.yml mentor'
    }
}