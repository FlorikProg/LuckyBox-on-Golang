document.getElementById('openButton').addEventListener('click', function() {
    var boxImage = document.getElementById('boxImage');
    var button = document.getElementById('openButton');
    
    // Remove the button immediately
    button.remove();
    
    // Add fade-out class to start the transition
    boxImage.classList.add('fade-out');

    // Wait for the fade-out transition to complete
    setTimeout(function() {
        // Generate a random number between 0 and 1
        var randomNumber = Math.random();

        // 80% chance to show "open.png", 20% chance to show a diamond emoji
        if (randomNumber < 0.8) {
            boxImage.src = 'C:\Users\Roman\Desktop\Web\src\file\img\open.png'; // Change to the path of the open box image
        } else {
            // Replace the image with a diamond emoji
            boxImage.style.display = 'none'; // Hide the image element
            var emoji = document.createElement('span');
            var text_win = document.createElement('h2');
            text_win.textContent = 'ВЫЙГРЫШ';
            emoji.textContent = '💎';
            emoji.className = 'diamond-emoji'; // Apply the diamond emoji class
            boxImage.parentElement.appendChild(emoji);
            boxImage.parentElement.appendChild(text_win);

            // Get the position of the diamond emoji
            var rect = emoji.getBoundingClientRect();
            var x = (rect.left + rect.right) / 2 / window.innerWidth;
            var y = (rect.top + rect.bottom) / 2 / window.innerHeight;

            // Trigger confetti effect from the diamond's position
            confetti({
                particleCount: 200, // Number of particles
                spread: 360, // Spread particles in all directions
                origin: { x: x, y: y } // Start from the diamond's position
            });
        }
        
        // Remove the fade-out class to fade the image back in
        boxImage.classList.remove('fade-out');
    }, 500); // Match this duration with the CSS transition duration
});
