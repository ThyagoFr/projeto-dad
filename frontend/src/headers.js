export default () => {
    return {
      Authorization: `Bearer ${localStorage.token}`,
      'Access-Control-Allow-Origin': '*',
    }
  }
  