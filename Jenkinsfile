pipeline {
  agent any
  stages {
    stage('init') {
      steps {
        tool(name: 'jdk-11', type: 'jdk')
        tool(name: 'maven-3', type: 'maven')
        sh '''env

java -version

mvn  -version'''
      }
    }

    stage('build') {
      steps {
        sh '''env

java -version

mvn  -version'''
      }
    }

    stage('archive') {
      steps {
        archiveArtifacts 'user-api-server/target/user-api-server-0.1.jar,mine-api-server/target/mine-api-server-0.1.jar,activity-api-server/target/activity-api-server-0.1.jar,im-api-server/target/im-api-server-0.1.jar,notification-api-server/target/notification-api-server-0.1.jar,blackhole-api-server/target/blackhole-api-server-0.1.jar'
      }
    }

    stage('dev') {
      parallel {
        stage('dev') {
          steps {
            sh 'ls'
          }
        }

        stage('test') {
          steps {
            sh 'ls'
          }
        }

        stage('stage') {
          steps {
            sh 'ls'
          }
        }

        stage('prrod') {
          steps {
            sh 'ls'
          }
        }

      }
    }

    stage('done') {
      steps {
        sh 'ls'
      }
    }

  }
}