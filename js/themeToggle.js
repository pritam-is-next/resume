(function () {
    console.log('Theme script loaded');
    
    // Test localStorage first
    console.log('Testing localStorage...');
    localStorage.setItem('test', 'value');
    console.log('localStorage test:', localStorage.getItem('test'));
    
    const htmlEl = document.documentElement;
    const themeToggle = document.getElementById('themeToggle');
    
    console.log('htmlEl:', htmlEl);
    console.log('themeToggle:', themeToggle);
    
    if (!themeToggle) {
        console.error('themeToggle element not found');
        return;
    }
    
    // Load saved theme from localStorage or use system preference
    const savedTheme = localStorage.getItem('theme');
    const systemPreference = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    const initialTheme = savedTheme || systemPreference;
    
    console.log('Saved theme:', savedTheme);
    console.log('System preference:', systemPreference);
    console.log('Initial theme:', initialTheme);
    
    // Set initial theme
    htmlEl.setAttribute('data-bs-theme', initialTheme);
    updateThemeToggleButton(themeToggle, initialTheme);
    
    // Handle theme toggle
    themeToggle.addEventListener('click', function () {
        const currentTheme = htmlEl.getAttribute('data-bs-theme');
        const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
        
        console.log('Switching from', currentTheme, 'to', newTheme);
        
        htmlEl.setAttribute('data-bs-theme', newTheme);
        updateThemeToggleButton(this, newTheme);
        localStorage.setItem('theme', newTheme);
        
        console.log('Theme saved to localStorage:', localStorage.getItem('theme'));
    });
})();

function updateThemeToggleButton(button, theme) {
    if (!button) return;
    button.innerHTML = theme === 'dark' ?
        '<i class="bi bi-sun-fill me-2"></i>Light Mode' :
        '<i class="bi bi-moon-stars-fill me-2"></i>Dark Mode';
}