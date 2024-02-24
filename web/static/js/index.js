document.addEventListener('DOMContentLoaded', function() {
  const form = document.getElementById('form-expression');
  const _ul = document.querySelector('.list-group');
  
  function serializeForm(formNode) {
    return new FormData(formNode)
  };
  
  async function sendData(data) {
    return await fetch('http://localhost:8081/expressions', {
      method: 'POST',
      mode: 'cors',
      body: data,
    })
  };

  async function handleFormSubmit(event) {
    let res;
    event.preventDefault();
    const data = serializeForm(form);
    data.append('request_id', crypto.randomUUID())
    
    if (form.elements.expression.value) {
      res = await sendData(data);
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
  
  form.addEventListener('submit', handleFormSubmit)

})


