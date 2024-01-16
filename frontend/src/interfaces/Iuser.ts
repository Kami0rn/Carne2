import { WalletInterface } from "./Iwallet";

export interface UserInterface {    

    FirstName?: string;
    LastName?: string
    PhoneNumber?: number;
    Email?: string;
    Password?: string;
    Wallet?: WalletInterface;


}