node {

    stage('pull changes') {
        git 'https://github.com/ottmartens/mentor-server'
    }
    
    stage('build Docker image, run tests') {
        sh 'cp /srv/mentor-server/.env .'
        sh 'docker build -t mentor-server .'
    }

    stage('push image to local registry') {
        sh 'docker tag mentor-server localhost:5000/mentor-server-local'
        sh 'docker push localhost:5000/mentor-server-local'
    }

    
    stage('deploy & reload service') {

        withCredentials([
            string(credentialsId: 'POSTGRES_USER', variable: 'POSTGRES_USER'),
            string(credentialsId: 'POSTGRES_PASSWORD', variable: 'POSTGRES_PASSWORD'),
            string(credentialsId: 'POSTGRES_DB', variable: 'POSTGRES_DB')]) {
            
            sh 'docker stack deploy -c docker-compose.yml mentor'
        }
    }

    stage('kill old containers') {
        sh 'docker container prune -f'
    }
}