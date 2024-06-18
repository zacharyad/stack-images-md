document.addEventListener('DOMContentLoaded', function () {
  let selectedLogos = [];
  let countSelected = 0;

  const logoGrid = document.getElementById('logo-grid');
  const generatedUrlElement = document.getElementById('generated-url');
  const copyUrlButton = document.getElementById('copy-url-button');
  const clearUrlButton = document.getElementById('clear-url-button');

  // Fetch the list of images from the server
  fetch('/images-list')
    .then((response) => response.json())
    .then((logos) => {
      logos.forEach((logo) => {
        const logoItem = document.createElement('div');
        const orderCount = document.createElement('div');

        logoItem.className = 'logo-item';
        logoItem.countFlag = 0;
        logoItem.logoName = logo;
        logoItem.classList.add('logoimage');

        orderCount.className = 'order-count-flag';
        orderCount.style.opacity = 0;

        logoItem.appendChild(orderCount);

        const logoImage = document.createElement('img');
        logoImage.src = `https://www.stackimages.xyz/l/${logo}`;
        logoImage.alt = logo;
        logoImage.style.width = '100%';
        logoImage.style.height = 'auto';

        logoItem.appendChild(logoImage);

        logoItem.addEventListener('click', () => {
          if (logoItem.classList.contains('selected')) {
            logoItem.classList.remove('selected');

            selectedLogos = selectedLogos.filter((item) => item !== logo);
            orderCount.style.opacity = 0;
            countSelected--;
            reorderAllFlags(logoItem.countFlag);
          } else {
            logoItem.classList.add('selected');
            selectedLogos.push(logo);
            orderCount.style.opacity = 100;
            clearUrlButton.style.opacity = 100;

            ++countSelected;
            logoItem.countFlag = countSelected;
            orderCount.innerText = countSelected;
          }

          updateUrl();

          if (generatedUrlElement.value.length === 30) {
            generatedUrlElement.value = '';
            copyUrlButton.disabled = true;
            clearUrlButton.disabled = true;
            clearUrlButton.style.opacity = 0;
          }
        });

        logoItem.addEventListener('touch', () => {
          if (logoItem.classList.contains('selected')) {
            logoItem.classList.remove('selected');
            selectedLogos = selectedLogos.filter((item) => item !== logo);
            orderCount.style.opacity = 0;
            countSelected--;
            reorderAllFlags(logoItem.countFlag);
          } else {
            logoItem.classList.add('selected');
            selectedLogos.push(logo);
            orderCount.style.opacity = 100;
            ++countSelected;
            logoItem.countFlag = countSelected;
            orderCount.innerText = countSelected;
          }

          updateUrl();

          if (generatedUrlElement.value.length === 30) {
            generatedUrlElement.value = '';
            copyUrlButton.disabled = true;
            clearUrlButton.disabled = true;
          }
        });

        logoGrid.appendChild(logoItem);
      });
    })
    .catch((error) => console.error('Error fetching images:', error));

  function updateUrl() {
    if (copyUrlButton.innerText === 'Copied') {
      copyUrlButton.innerText = 'Copy URL';
      copyUrlButton.style.backgroundColor = '#007aff';
    }
    if (clearUrlButton.innerText === 'Cleared') {
      clearUrlButton.innerText = 'Clear Selections';
      clearUrlButton.style.backgroundColor = '#ff948b';
    }

    copyUrlButton.disabled = false;
    clearUrlButton.disabled = false;

    const url = `https://www.stackimages.xyz/l/${selectedLogos
      .map((item) => item.split('/').pop().split('.')[0])
      .join('-')}`;

    generatedUrlElement.value = url;
  }

  // eventListerns

  clearUrlButton.addEventListener('click', () => {
    generatedUrlElement.value = '';

    clearAllSelections();
  });

  copyUrlButton.addEventListener('click', () => {
    const url = generatedUrlElement.value;

    navigator.clipboard
      .writeText(url)
      .then(() => {
        copyUrlButton.innerText = 'Copied';
        copyUrlButton.style.backgroundColor = '#49D26D';
      })
      .catch((err) => {
        console.error('Failed to copy URL: ', err);
      });
  });

  function reorderAllFlags(numberDeleted) {
    console.log('numberDeleted: ', numberDeleted);

    let allLogos = [...document.querySelectorAll('.logoimage')]
      .filter((elem) => {
        return selectedLogos.includes(elem.logoName);
      })
      .map((elem) => {
        return {
          parent: elem,
          flag: elem.children[0],
          prevCount: elem.countFlag - 1,
        };
      });

    countSelected = allLogos.length;

    allLogos.forEach((elem) => {
      if (elem.prevCount >= numberDeleted) {
        elem.flag.innerText = elem.prevCount;
        elem.parent.countFlag = elem.prevCount;
      }
    });
  }

  function clearAllSelections() {
    let allLogos = [...document.querySelectorAll('.logoimage')]
      .filter((elem) => {
        return selectedLogos.includes(elem.logoName);
      })
      .forEach((elem) => {
        resetFlagsOnSelectedElems(elem);
      });

    resetGlobalVars();

    clearUrlButton.classList.toggle('fade');
    clearUrlButton.style.opacity = 0;
    copyUrlButton.disabled = true;
  }

  function resetGlobalVars() {
    selectedLogos = [];
    countSelected = selectedLogos.length;
  }

  function resetFlagsOnSelectedElems(elem) {
    let flagElem = elem.children[0];

    elem.classList.remove('selected');
    flagElem.innerText = 0;
    flagElem.style.opacity = 0;
    generatedUrlElement.disabled = true;
    clearUrlButton.innerText = 'Cleared';
    elem.countFlag = 0;
  }
});
