const form = document.getElementById('form-expression');

function serializeForm(formNode) {
  return new FormData(formNode)
};

async function sendData(data) {
  return await fetch('/getexp', {
    method: 'POST',
    headers: { 'Content-Type': 'multipart/form-data' },
    body: data,
  })
};

async function handleFormSubmit(event) {
  event.preventDefault();
  const ul = document.querySelector('.list-group');
  const input = document.querySelector('#input-expression');
  const exp = input.value;
  if (exp != '') {
    const list = document.createElement('li');
    list.classList.add('list-group-item');
    list.textContent = exp;
    ul.append(list);
    // input.value = '';
  }

  const data = serializeForm(event.target)
  console.log(data);
  const { status, error } = await sendData(data)

  if (status == 200) {
    console.log('Данные успешно отправлены!')
  } else (
    console.log(error)
  )
};

form.addEventListener('submit', handleFormSubmit)
