document.addEventListener('DOMContentLoaded', () => {
  const formRegister = document.forms.formRegister;

  async function handleFormRegisterSubmit(event) {
    let res;
    event.preventDefault();

    if (formRegister.elements.name.value &&
      formRegister.elements.email.value &&
      formRegister.elements.password.value) {
      const name = formRegister.elements.name.value;
      const email = formRegister.elements.email.value;
      const password = formRegister.elements.password.value;

      res = await fetch('http://localhost:8081/api/v1/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Headers': 'Authorization',
        },
        body: JSON.stringify({
          name,
          email,
          password
        }),
      });
    }

    const response = await res.json();
    if (res.ok) {
      const token = "Bearer " + response.token
      sessionStorage.setItem('token', token);
      location.replace('/')
    }
  };

  if (formRegister) {
    formRegister.addEventListener('submit', handleFormRegisterSubmit);
  }
})