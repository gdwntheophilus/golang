import { UrlsService } from 'src/app/url/urls.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-viewcontact',
  templateUrl: './viewcontact.component.html',
  styleUrls: ['./viewcontact.component.css'],
})
export class ViewcontactComponent implements OnInit {
  payload: any = {
    emailid: '',
  };

  contacts: any = {};

  constructor(private http: HttpClient, private urlsService: UrlsService) {}

  ngOnInit(): void {
    this.payload.emailid = sessionStorage.getItem('token');
    const headers = new HttpHeaders();
    // var headers = { 'content-type': 'application/json' };
    headers.set('Content-Type', 'application/json; charset=utf-8');

    console.log(this.payload);
    console.log(this.urlsService.urls().getAllContacts);
    this.http
      .post(this.urlsService.urls().getAllContacts, this.payload, { headers })
      .subscribe((response) => {
        this.contacts = response;
        console.log(response);
      });
  }
}
