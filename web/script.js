document.addEventListener('DOMContentLoaded', function() {
  const randomQuoteDiv = document.getElementById('random-quote');
  const refreshRandomButton = document.getElementById('refresh-random');
  const quoteIdInput = document.getElementById('quote-id-input');
  const searchButton = document.getElementById('search-button');
  const searchResultDiv = document.getElementById('search-result');
  const quotesList = document.getElementById('quotes-list');
  const prevPageButton = document.getElementById('prev-page');
  const nextPageButton = document.getElementById('next-page');
  const currentPageSpan = document.getElementById('current-page');

  let currentPage = 1;

  // Rastgele alıntıyı API'den yükle.
  function loadRandomQuote() {
    // 1 ile 1000 arasında rastgele bir ID üretiyoruz.
    const randomID = Math.floor(Math.random() * 1000) + 1;
    fetch(`http://localhost:5000/quotes/${randomID}`)
      .then(response => {
        if (!response.ok) {
          throw new Error('Quote not found');
        }
        return response.json();
      })
      .then(data => {
        randomQuoteDiv.textContent = `${data.author}: "${data.text}"`;
      })
      .catch(error => {
        randomQuoteDiv.textContent = 'Rastgele alıntı yüklenirken hata oluştu.';
        console.error(error);
      });
  }

  // Belirtilen ID'ye göre alıntıyı arar.
  function searchQuoteByID() {
    const id = quoteIdInput.value;
    if (!id) {
      searchResultDiv.textContent = 'Lütfen bir ID girin.';
      return;
    }
    fetch(`http://localhost:5000/quotes/${id}`)
      .then(response => {
        if (!response.ok) {
          throw new Error('Quote not found');
        }
        return response.json();
      })
      .then(data => {
        searchResultDiv.textContent = `${data.author}: "${data.text}"`;
      })
      .catch(error => {
        searchResultDiv.textContent = 'Quote bulunamadı.';
        console.error(error);
      });
  }

  // Belirtilen sayfa numarasına göre alıntıları yükler.
  function loadQuotesPage(page) {
    fetch(`http://localhost:5000/quotes?page=${page}`)
      .then(response => response.json())
      .then(data => {
        quotesList.innerHTML = '';
        if (data.length === 0) {
          quotesList.innerHTML = '<li>Bu sayfada alıntı bulunamadı.</li>';
          return;
        }
        data.forEach(quote => {
          const li = document.createElement('li');
          li.textContent = `${quote.id} - ${quote.author}: "${quote.text}"`;
          quotesList.appendChild(li);
        });
      })
      .catch(error => {
        quotesList.innerHTML = '<li>Alıntılar yüklenirken hata oluştu.</li>';
        console.error(error);
      });
  }

  // Event Listener'lar
  refreshRandomButton.addEventListener('click', loadRandomQuote);
  searchButton.addEventListener('click', searchQuoteByID);
  prevPageButton.addEventListener('click', function() {
    if (currentPage > 1) {
      currentPage--;
      currentPageSpan.textContent = currentPage;
      loadQuotesPage(currentPage);
    }
  });
  nextPageButton.addEventListener('click', function() {
    currentPage++;
    currentPageSpan.textContent = currentPage;
    loadQuotesPage(currentPage);
  });

  // İlk yüklemeler
  loadRandomQuote();
  loadQuotesPage(currentPage);
});
