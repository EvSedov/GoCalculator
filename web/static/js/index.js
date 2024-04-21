document.addEventListener('DOMContentLoaded', function () {
  const formExpression = document.getElementById('form-expression');
  const formRegister = document.forms.formRegister
  const _ul = document.querySelector('.list-group');

  function serializeForm(formNode) {
    return new FormData(formNode)
  };

  async function sendData(url, data) {
    return await fetch(url, {
      method: 'POST',
      mode: 'cors',
      body: data,
    })
  };

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
      sessionStorage.setItem('token', response.token);
      const auth = res.headers.get('Authorization')
      sessionStorage.setItem('auth', auth);
      location.replace('/login')
    }
  };

  async function handleFormExpressionSubmit(event) {
    let res;
    event.preventDefault();
    const dataExpression = serializeForm(formExpression);
    dataExpression.append('request_id', crypto.randomUUID())

    if (formExpression.elements.expression.value) {
      res = await sendData('http://localhost:8081/api/v1/expressions', dataExpression);
    }

    const response = await res.json();

    if (res.ok) {
      form.elements.expression.value = '';
      const _li = document.createElement('li')
      const _divWrap = document.createElement('div')
      const _div = document.createElement('div')
      const _span = document.createElement('span');
      _li.classList.add(...['list-group-item', 'd-flex', 'justify-content-between', 'align-items-start']);
      _divWrap.classList.add('ms-2', 'me-auto');
      _div.classList.add('fw-bold');
      _span.classList.add('badge', 'bg-primary', 'rounded-pill')
      _span.append(response.state)
      _div.append(response.expression);
      _divWrap.append(_div, response.message);
      _li.append(_divWrap, _span);
      _ul.append(_li);
    }
  };


  formRegister.addEventListener('submit', handleFormRegisterSubmit);
  if (formExpression) {
    formExpression.addEventListener('submit', handleFormExpressionSubmit);
  }
})


