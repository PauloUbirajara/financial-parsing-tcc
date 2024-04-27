export function getFilteredUrlSearchParams(
  object: Record<any, any>,
): URLSearchParams {
  const urlParams = new URLSearchParams(
    Object.fromEntries(
      Object.entries(object).filter(
        (o) => ![null, undefined, "null", "undefined"].includes(o[1]),
      ),
    ),
  );
  return urlParams;
}
