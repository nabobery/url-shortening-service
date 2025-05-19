document.addEventListener("DOMContentLoaded", () => {
  const longUrlInput = document.getElementById("longUrl");
  const shortenBtn = document.getElementById("shortenBtn");
  const resultSection = document.getElementById("resultSection");
  const shortUrlInput = document.getElementById("shortUrl");
  const copyBtn = document.getElementById("copyBtn");
  const createdAtSpan = document.getElementById("createdAt");
  const visitCountSpan = document.getElementById("visitCount");
  const errorSection = document.getElementById("errorSection");
  const errorMessage = document.getElementById("errorMessage");

  // Base URL for the API
  const apiBaseUrl = window.location.origin;

  // Function to show error message
  const showError = (message) => {
    errorMessage.textContent = message;
    errorSection.classList.remove("hidden");
    resultSection.classList.add("hidden");
  };

  // Function to hide error message
  const hideError = () => {
    errorSection.classList.add("hidden");
  };

  // Function to format date
  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleString();
  };

  // Function to shorten URL
  const shortenUrl = async (url) => {
    try {
      const response = await fetch(`${apiBaseUrl}/api/shorten`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ url }),
      });

      if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || "Failed to shorten URL");
      }

      const data = await response.json();
      return data;
    } catch (error) {
      throw new Error(error.message || "Failed to shorten URL");
    }
  };

  // Function to get URL stats
  const getUrlStats = async (shortCode) => {
    try {
      const response = await fetch(`${apiBaseUrl}/api/urls/${shortCode}`);
      if (!response.ok) {
        throw new Error("Failed to get URL statistics");
      }
      return await response.json();
    } catch (error) {
      console.error("Error fetching stats:", error);
      return null;
    }
  };

  // Handle form submission
  shortenBtn.addEventListener("click", async () => {
    const url = longUrlInput.value.trim();

    if (!url) {
      showError("Please enter a URL");
      return;
    }

    try {
      hideError();
      shortenBtn.disabled = true;
      shortenBtn.textContent = "Shortening...";

      const result = await shortenUrl(url);
      const shortUrl = `${apiBaseUrl}/${result.shortCode}`;

      // Get stats for the newly created URL
      const stats = await getUrlStats(result.shortCode);

      // Update UI
      shortUrlInput.value = shortUrl;
      if (stats) {
        createdAtSpan.textContent = formatDate(stats.createdAt);
        visitCountSpan.textContent = stats.accessCount;
      }

      resultSection.classList.remove("hidden");
    } catch (error) {
      showError(error.message);
    } finally {
      shortenBtn.disabled = false;
      shortenBtn.textContent = "Shorten URL";
    }
  });

  // Handle copy button
  copyBtn.addEventListener("click", async () => {
    try {
      await navigator.clipboard.writeText(shortUrlInput.value);
      const originalText = copyBtn.innerHTML;
      copyBtn.innerHTML = '<i class="fas fa-check"></i>';
      setTimeout(() => {
        copyBtn.innerHTML = originalText;
      }, 2000);
    } catch (err) {
      showError("Failed to copy URL");
    }
  });

  // Handle Enter key in input
  longUrlInput.addEventListener("keypress", (e) => {
    if (e.key === "Enter") {
      shortenBtn.click();
    }
  });
});
