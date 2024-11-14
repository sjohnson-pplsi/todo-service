import { useMemo } from "react";
import { useForm, UseFormProps, UseFormReturn } from "react-hook-form";

import { yupResolver } from "@hookform/resolvers/yup";
import { AnyObjectSchema } from "yup";

export type UseYupFormReturn<
  TSchema extends Record<string, unknown>,
  TContext extends Record<string, unknown> = Record<string, unknown>
> = UseFormReturn<TSchema, TContext>;

export const useYupForm = <
  TSchema extends AnyObjectSchema,
  TContext extends Record<string, unknown> = Record<string, unknown>
>(
  schema: TSchema,
  {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    defaultValues = {} as any,
    ...options
  }: UseFormProps<ReturnType<TSchema["validateSync"]>, TContext> = {}
) => {
  const schemaDefault = schema.getDefault();
  const merged = useMemo(() => {
    return deepMerge(defaultValues, schemaDefault);
  }, [defaultValues, schemaDefault]);
  return useForm<ReturnType<TSchema["validateSync"]>, TContext>({
    defaultValues: merged,
    resolver: yupResolver(schema),
    ...options,
  });
};

function deepMerge<TA, TB>(valueA: TA, valueB: TB): TA & TB {
  switch (typeof valueA) {
    case "object":
      if (valueA === null) {
        // @ts-expect-error TODO: Fix this type
        return valueB;
      }
      switch (typeof valueB) {
        case "undefined":
          // @ts-expect-error TODO: Fix this type
          return valueA;
        case "object":
          if (valueB === null) {
            // @ts-expect-error TODO: Fix this type
            return valueA;
          }
          // @ts-expect-error TODO: Fix types
          return mergeObjects(valueA, valueB);
        default:
          // @ts-expect-error TODO: Fix this type
          return valueB;
      }
    case "undefined":
      // @ts-expect-error TODO: Fix this type
      return valueB;
    default:
      // @ts-expect-error TODO: Fix this type
      return valueA || valueB;
  }
}

function mergeObjects<
  TA extends Record<string, unknown>,
  TB extends Record<string, unknown>
>(valueA: TA, valueB: TB): TA & TB {
  const keysHash: Record<string, boolean> = {};
  // @ts-expect-error TODO: Fix this type
  const mergeHash: TA & TB = {};
  Object.keys(valueA).forEach((key) => (keysHash[key] = true));
  Object.keys(valueB).forEach((key) => (keysHash[key] = true));
  Object.keys(keysHash).forEach((key) => {
    // @ts-expect-error TODO: Fix this type
    mergeHash[key as keyof TA & TB] = deepMerge(
      valueA[key as keyof TA],
      valueB[key as keyof TB]
    );
  });
  return mergeHash;
}
