document.addEventListener('DOMContentLoaded', () => {
  var popoverTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]'))
  var popoverList = popoverTriggerList.map(function (popoverTriggerEl) {
    return new bootstrap.Popover(popoverTriggerEl)
  })

  const logout = document.getElementById('logout');

  logout.addEventListener('click', (event) => {
    event.preventDefault()
    sessionStorage.clear();
  });
});
