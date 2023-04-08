CREATE TABLE search_history (
  id INT AUTO_INCREMENT PRIMARY KEY,
  web VARCHAR(30),
  category VARCHAR(30),
  search_results JSON
);
