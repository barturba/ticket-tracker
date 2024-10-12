// Users

export async function fetchUsers() {
  try {
    const url = new URL(`http://localhost:8080/v1/users`);

    const searchParams = url.searchParams;
    searchParams.set("sort", "last_name");
    searchParams.set("page", "1");
    searchParams.set("page_size", "100");
    const data = await fetch(url.toString(), {
      method: "GET",
    });
    // Simulate slow load
    // await new Promise((resolve) => setTimeout(resolve, 1000));
    if (data.ok) {
      const users = await data.json();
      if (users) {
        return users;
      } else {
        return [];
      }
    }
  } catch (error) {
    console.log(`fetchUsers error: ${error}`);
    throw new Error("Failed to fetch users data.");
  }
}
