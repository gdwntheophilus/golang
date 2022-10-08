import { Router } from '@angular/router';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-adressbook',
  templateUrl: './adressbook.component.html',
  styleUrls: ['./adressbook.component.css'],
  encapsulation: ViewEncapsulation.None,
})
export class AdressbookComponent implements OnInit {
  constructor(private router: Router) {
    if (sessionStorage.getItem('token')) {
    } else {
      router.navigate(['login']);
    }
  }

  showAddContact = true;
  showViewContact = false;

  ngOnInit(): void {}

  handleAddContact() {
    this.showAddContact = true;
    this.showViewContact = false;
  }

  handleViewContact() {
    this.showAddContact = false;
    this.showViewContact = true;
  }
}
