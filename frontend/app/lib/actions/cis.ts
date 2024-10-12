export async function fetchCIs() {
  try {
    const url = new URL(`http://localhost:8080/v1/cis`);
    const searchParams = url.searchParams;
    searchParams.set("sort", "name");
    searchParams.set("page", "1");
    searchParams.set("page_size", "100");
    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const configurationItems = await data.json();
      if (configurationItems) {
        return configurationItems;
      } else {
        return [];
      }
    }
  } catch (error) {
    console.log(`fetchCIs error: ${error}`);
    throw new Error("Failed to fetch configuration items data.");
  }
}
