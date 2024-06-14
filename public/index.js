document.addEventListener('DOMContentLoaded', function () {
  const logoGrid = document.getElementById('logo-grid');
  const generatedUrlElement = document.getElementById('generated-url');
  const copyUrlButton = document.getElementById('copy-url-button');

  let selectedLogos = [];

  // Fetch the list of images from the server
  fetch('/images-list')
    .then((response) => response.json())
    .then((logos) => {
      logos.forEach((logo) => {
        const logoItem = document.createElement('div');
        logoItem.className = 'logo-item';

        const logoImage = document.createElement('img');
        logoImage.src = `https://www.stackimages.xyz/l${logo}`;
        logoImage.alt = logo;
        logoImage.style.width = '100%';
        logoImage.style.height = 'auto';

        logoItem.appendChild(logoImage);

        logoItem.addEventListener('click', () => {
          if (logoItem.classList.contains('selected')) {
            logoItem.classList.remove('selected');
            selectedLogos = selectedLogos.filter((item) => item !== logo);
          } else {
            logoItem.classList.add('selected');
            selectedLogos.push(logo);
          }
          updateUrl();
        });

        logoGrid.appendChild(logoItem);
      });
    })
    .catch((error) => console.error('Error fetching images:', error));

  function updateUrl() {
    const url = `https://www.stackimages.xyz/l/${selectedLogos
      .map((item) => item.split('/').pop().split('.')[0])
      .join('-')}`;
    generatedUrlElement.innerText = url;
  }

  copyUrlButton.addEventListener('click', () => {
    const url = generatedUrlElement.innerText;
    navigator.clipboard
      .writeText(url)
      .then(() => {
        alert('URL copied to clipboard');
      })
      .catch((err) => {
        console.error('Failed to copy URL: ', err);
      });
  });
});
