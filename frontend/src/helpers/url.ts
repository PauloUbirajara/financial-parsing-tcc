export function getFilteredUrlSearchParams(
  object: Record<any, any>,
): URLSearchParams {
  return new URLSearchParams(
    Object.fromEntries(
      Object.entries(object).filter((o) => ![null, undefined].includes(o[1])),
    ),
  );
}
