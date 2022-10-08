import { Router } from '@angular/router';
import { UrlsService } from './../../url/urls.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-addcontact',
  templateUrl: './addcontact.component.html',
  styleUrls: ['./addcontact.component.css'],
})
export class AddcontactComponent implements OnInit {
  loginStatus = '';
  user: any = {
    emailid: '',
    contacttype: '',
    name: '',
    relationship: '',
  };
  constructor(
    private http: HttpClient,
    private urlsService: UrlsService,
    private router: Router
  ) {}

  ngOnInit(): void {}
  UserLoginSubmit() {
    console.log(this.user);
    if (this.user.phonenumber) {
      if (
        this.user.name === '' ||
        this.user.contacttype === '' ||
        this.user.relationship === ''
        // this.user.emailid === ''
      ) {
        this.loginStatus = 'Kindly fill in the blank details!!';
      } else {
        var headers = new HttpHeaders({
          'Content-Type': 'application/json',
        });
        console.log(this.user);
        this.user.emailid = sessionStorage.getItem('token');
        this.http
          .post(this.urlsService.urls().saveContact, [this.user], {
            headers,
          })
          .subscribe((response: any) => {
            this.loginStatus = response.message;
            this.user = {
              emailid: '',
              contacttype: '',
              name: '',
              relationship: '',
            };
          });
        this.router.navigate(['/address']);
      }
    } else {
      this.loginStatus = 'Kindly provide Phone Details!!';
    }
  }
}
