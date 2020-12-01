export function ucFirst(str) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

export function toUpperCase(value) {
  let v = value.trim();
  if (v !== '') {
    return v.toUpperCase();
  } else {
    return v;
  }
}
