function postForm(event) {
    event.preventDefault(); 
  
    const taskName = document.getElementById('taskName').value;
    const taskDescription = document.getElementById('taskDesc').value;
    const currentDate = new Date().toISOString().split('T')[0]; 
  
    const data = {
      title: taskName,
      description: taskDescription,
      date: currentDate
    };
  
    fetch('http://127.0.0.1:8080/create-task/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
      window.location.assign("file:///E:/others/SW/Projects/Go-LP/Golang-Learning-phase/Task-Manager/view/index.html")
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  }
  document.getElementById('taskForm').addEventListener('submit', postForm);
  

