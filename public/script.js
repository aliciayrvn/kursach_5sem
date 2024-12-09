document.getElementById('register-form').addEventListener('submit', function(event) {
    event.preventDefault();
  
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
  
    // Отправка данных на сервер
    fetch('http://localhost:8081/users/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password })
    })
    .then(response => response.json())
    .then(data => {
      if (data.error) {
        document.getElementById('message').textContent = `Error: ${data.error}`;
      } else {
        document.getElementById('message').textContent = `User registered successfully!`;
      }
    })
    .catch(error => {
      document.getElementById('message').textContent = `Error: ${error}`;
    });
  });
  
  document.getElementById('booking-form').addEventListener('submit', function(event) {
    event.preventDefault();
  
    const user_id = document.getElementById('user_id').value;
    const car_id = document.getElementById('car_id').value;
    const start_date = document.getElementById('start_date').value;
    const end_date = document.getElementById('end_date').value;
  
    // Отправка данных для бронирования
    fetch(`http://localhost:8081/bookings/${car_id}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ user_id, car_id, start_date, end_date })
    })
    .then(response => response.json())
    .then(data => {
      if (data.error) {
        document.getElementById('message').textContent = `Error: ${data.error}`;
      } else {
        document.getElementById('message').textContent = `Booking created successfully!`;
      }
    })
    .catch(error => {
      document.getElementById('message').textContent = `Error: ${error}`;
    });
  });
  