if (!sessionStorage.token) {
  location.replace('/login');
}

const url = 'http://localhost:8081/api/v1/expressions';

async function loadExpressions() {
  const _ul = document.querySelector('.list-group');
  const res = await fetch(url, {
    method: 'GET',
    mode: 'cors',
    headers: {
      'Authorization': sessionStorage.token,
    },
  });

  const response = await res.json();
  if (response.message == "ok") {
    _ul.innerHTML = '';
    const userExpressions = [...response["user_expressions"]];
    if (userExpressions.length > 0) {
      userExpressions.forEach((item) => {
        const _li = document.createElement('li')
        const _divWrap = document.createElement('div')
        const _div = document.createElement('div')
        const _span = document.createElement('span');
        _li.classList.add(...['list-group-item', 'd-flex', 'justify-content-between', 'align-items-start']);
        _divWrap.classList.add('ms-2', 'me-auto');
        _div.classList.add('fw-bold');
        _span.classList.add('badge', 'bg-primary', 'rounded-pill')
        _span.append(item.state)
        _div.append(item.expression);
        _divWrap.append(_div, item.message);
        _li.append(_divWrap, _span);
        _ul.append(_li);
      });
    } else {
      const _li = document.createElement('li');
      const _divWrap = document.createElement('div');
      const _div = document.createElement('div');
      _li.classList.add(...['list-group-item', 'd-flex', 'justify-content-between', 'align-items-start']);
      _divWrap.classList.add('ms-2', 'me-auto');
      _div.classList.add('fw-bold');
      _div.append("Выражения на сервер не отправлялись");
      _divWrap.append(_div, 'Для отправки выражения введите его в поле для ввода и нажмите кнопку "Send"');
      _li.append(_divWrap)
      _ul.append(_li);
    }
  }
}

async function handleFormExpressionSubmit(event) {
  let res;
  const _ul = document.querySelector('.list-group');
  event.preventDefault();
  if (formExpression.elements.expression.value) {
    const expression = formExpression.elements.expression.value
    const request_id = crypto.randomUUID()
    res = await fetch(url, {
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

document.addEventListener('DOMContentLoaded', function () {
  const formExpression = document.forms.formExpression;
  loadExpressions();
  if (formExpression) {
    formExpression.addEventListener('submit', handleFormExpressionSubmit);
  }
})


