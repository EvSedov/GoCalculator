document.addEventListener('DOMContentLoaded', () => {
  const formLogin = document.forms.formLogin;

  async function handleFormRegisterSubmit(event) {
    let res;
    event.preventDefault();

    if (formLogin.elements.email.value &&
      formLogin.elements.password.value) {
      const email = formLogin.elements.email.value;
      const password = formLogin.elements.password.value;

      res = await fetch('http://localhost:8081/api/v1/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Headers': 'Authorization',
        },
        body: JSON.stringify({
          email,
          password
        }),
      });
    }

    const response = await res.json();
    if (res.ok) {
      sessionStorage.setItem('token', response.token);
      const auth = res.headers.get('Authorization')
      sessionStorage.setItem('auth', auth);
      location.replace('/')
    }
  };

  if (formLogin) {
    formLogin.addEventListener('submit', handleFormRegisterSubmit);
  }
});
