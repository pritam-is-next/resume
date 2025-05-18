document.addEventListener('DOMContentLoaded', () => {
  const professions = ['Designer', 'Developer', 'Freelancer', 'Photographer'];
  const element = document.getElementById('profession');
  const cursor = document.querySelector('.typing-cursor');
  let index = 0;
  let charIndex = 0;
  let isDeleting = false;

  function type() {
    const current = professions[index];
    
    if (!isDeleting) {
      // Typing phase
      element.textContent = current.slice(0, charIndex + 1);
      charIndex++;
      
      if (charIndex === current.length) {
        isDeleting = true;
        setTimeout(type, 2000);
        return;
      }
    } else {
      // Deleting phase
      element.textContent = current.slice(0, charIndex - 1);
      charIndex--;
      
      if (charIndex === 0) {
        isDeleting = false;
        index = (index + 1) % professions.length;
      }
    }

    // Adjust speed for typing vs deleting
    const speed = isDeleting ? 75 : 150;
    setTimeout(type, speed);
  }

  // Start animation after title appears
  setTimeout(type, 1500);
});