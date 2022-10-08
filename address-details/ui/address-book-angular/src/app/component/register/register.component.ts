import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from './../../user/user';
import { Component, OnInit } from '@angular/core';
import { UrlsService } from 'src/app/url/urls.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
})
export class RegisterComponent implements OnInit {
  user: any = {
    emailid: '',
    password: '',
    retypepassword: '',
    firstname: '',
    lastname: '',
  };
  registerStatus = '';
  isRegistered = false;
  urlsService: any;

  constructor(private http: HttpClient) {
    this.urlsService = new UrlsService();
  }

  ngOnInit(): void {}
  UserRegisterSubmit() {
    console.log(this.user);
    if (
      this.user.emailid === '' ||
      this.user.password === '' ||
      this.user.retypepassword === '' ||
      this.user.firstname === '' ||
      this.user.lastname === ''
    ) {
      this.registerStatus =
        'User Details are empty, please provide proper details!!';
    } else {
      if (this.user.password === this.user.retypepassword) {
        var headers = new HttpHeaders({
          'Content-Type': 'application/json',
        });
        this.http
          .post(this.urlsService.urls().registerUser, this.user, {
            headers,
          })
          .subscribe((response: any) => {
            console.log(response);
            this.registerStatus = 'User Registered Successfully!!';
            this.isRegistered = true;
          });
      } else {
        this.registerStatus = 'Passwords do not match, please try again';
      }
    }
  }
}
