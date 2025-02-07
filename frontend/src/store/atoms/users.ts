import { atom } from "recoil";

export interface User {
  id: number;
  full_name: string;
  phone_number: string;
  latitude: number;
  longitude: number;
  created_at: string;
}

export const userState = atom<User[]>({
  key: "userState", // unique ID (with respect to other atoms/selectors)
  default: [], // default value (aka initial value)
});
