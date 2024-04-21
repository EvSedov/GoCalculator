if (!sessionStorage.token) {
  location.replace('/login');
}

document.addEventListener('DOMContentLoaded', function () {
  const formExpression = document.getElementById('form-expression');
  const _ul = document.querySelector('.list-group');

  async function handleFormExpressionSubmit(event) {
    let res;
    event.preventDefault();
    if (formExpression.elements.expression.value) {
      const expression = formExpression.elements.expression.value
      const request_id = crypto.randomUUID()

      res = await fetch('http://localhost:8081/api/v1/expressions', {
        method: 'POST',
        mode: 'cors',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': sessionStorage.token,
        },
        body: JSON.stringify({
          expression,
          request_id,
        }),
      });
    }

    const response = await res.json();

    if (res.ok) {
      formExpression.elements.expression.value = '';
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


