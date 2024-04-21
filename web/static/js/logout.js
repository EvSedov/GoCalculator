document.addEventListener('DOMContentLoaded', () => {
  const logout = document.getElementById('logout');

  logout.addEventListener('click', (event) => {
    event.preventDefault()
    sessionStorage.clear();
  });
});
