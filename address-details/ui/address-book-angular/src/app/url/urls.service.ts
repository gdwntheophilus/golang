import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class UrlsService {
  baseUrl = 'http://ec2-3-110-147-16.ap-south-1.compute.amazonaws.com:8080';
  public urls() {
    return {
      getRegisteredUser: this.baseUrl + '/user/getRegisteredUser',
      registerUser: this.baseUrl + '/user/register',
      saveContact: this.baseUrl + '/contact/saveContact',
      getAllContacts: this.baseUrl + '/contact/getAllContacts',
    };
  }

  constructor() {}
}
