import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class UrlsService {
  public urls() {
    return {
      getRegisteredUser: 'http://localhost:8080/user/getRegisteredUser',
    };
  }

  constructor() {}
}
