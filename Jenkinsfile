
​pipeline {
  agent any

  stages {
    stage('Git checkout') {
      steps {
        shell  git clone 'https://github.com/cloudflavor/dweller.git'
      }
    stage('Build') {
      steps {
          shell dep ensure
          shell make compile
      }
    }
    stage('Test') {
        steps {
          shell make test
        }
      }
    }
  }
}
​
