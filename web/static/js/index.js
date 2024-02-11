document.addEventListener('DOMContentLoaded', function() {
  const form = document.getElementById('form-expression');
  
  function serializeForm(formNode) {
    return new FormData(formNode)
  };
  
  async function sendData(data) {
    return await fetch('http://127.0.0.1:81/getexp', {
      method: 'POST',
      mode: 'cors',
      headers: { 'Content-Type': 'multipart/form-data' },
      body: data,
    })
  };
  
  async function handleFormSubmit(event) {
    event.preventDefault();

    const data = serializeForm(form);
    console.log([...data.entries()]);
    const res = await sendData(data);
    console.log(res)

    if (res.ok) {
      const ul = document.querySelector('.list-group');
      const input = document.querySelector('#input-expression');
      if (input.value !== '') {
        const list = document.createElement('li');
        list.classList.add('list-group-item');
        list.textContent = input.value;
        ul.append(list);
        input.value = '';
      }
    }
    
  };
  
  form.addEventListener('submit', handleFormSubmit)

})


