/*
Copyright (c) 2013 Landeshauptstadt München
Author Matthias S. Benkmann

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, 
MA  02110-1301, USA.
*/

package message

import (
         "../xml"
         "../config"
       )

// Sends "set_activated_for_installation" to client_addr and calls
// Send_new_ldap_config(client_addr, system)
func Send_set_activated_for_installation(client_addr string, system *xml.Hash) {
  // gosa-si-server sends LDAP config both before and after set_activated_for_installation
  // Personally I think that sending it BEFORE should be enough. But it doesn't hurt
  // to do it twice. Better safe than sorry.
  Send_new_ldap_config(client_addr, system)
  set_activated_for_installation := "<xml><header>set_activated_for_installation</header><set_activated_for_installation></set_activated_for_installation><source>"+ config.ServerSourceAddress +"</source><target>"+ client_addr +"</target></xml>"
  Client(client_addr).Tell(set_activated_for_installation, config.LocalClientMessageTTL)
  Send_new_ldap_config(client_addr, system)
}