if (!sessionStorage.auth) {
  location.replace('/login');
}

document.addEventListener('DOMContentLoaded', function () {
  const formExpression = document.getElementById('form-expression');
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

  if (formExpression) {
    formExpression.addEventListener('submit', handleFormExpressionSubmit);
  }
})


