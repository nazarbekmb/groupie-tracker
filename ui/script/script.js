  const triggerButton = document.querySelector('.trigger');
  const backdrop = document.querySelector('.backdrop');
  const fullScreenNav = document.querySelector('.full-screen');

  triggerButton.addEventListener('click', function() {
    backdrop.classList.add('open');
    fullScreenNav.classList.add('open');
  });
