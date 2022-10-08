import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class UrlsService {
  public urls() {
    return {
      getRegisteredUser: 'http://localhost:8080/user/getRegisteredUser',
      registerUser: 'http://localhost:8080/user/register',
      saveContact: 'http://localhost:8080/contact/saveContact',
      getAllContacts: 'http://localhost:8080/contact/getAllContacts',
    };
  }

  constructor() {}
}
