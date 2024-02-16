document.addEventListener('DOMContentLoaded', function() {
  const form = document.getElementById('form-expression');
  
  function serializeForm(formNode) {
    return new FormData(formNode)
  };
  
  async function sendData(data) {
    return await fetch('http://127.0.0.1:81/expressions', {
      method: 'POST',
      mode: 'cors',
      body: data,
    })
  };
  
  async function handleFormSubmit(event) {
    event.preventDefault();
    const data = serializeForm(form);
    data.append('request_id', crypto.randomUUID())
    const res = await sendData(data);
    if (res.ok) {
      response = await res.json()
      console.log(response)
      const ul = document.querySelector('.list-group');
      const input = document.querySelector('#input-expression');

      if (input.value !== '') {
        const list = document.createElement('li');
        const div = document.createElement('div');
        list.classList.add('list-group-item');
        div.classList.add('fw-bold');
        div.append(response.expression)
        list.append(div);
        list.append(response.message)
        ul.append(list);
        input.value = '';
      }
    }
  };
  
  form.addEventListener('submit', handleFormSubmit)

})


